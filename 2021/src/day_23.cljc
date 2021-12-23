(ns day-23 (:require util [clojure.pprint :refer [pprint]]))
(def filename "real-part-a")

(def input
  (apply vector (util/read-input #(apply vector (seq %)))))

(def depth (- (count input) 3))

(def direct-blocks
  (apply merge (for [from (range 1 5)
                     to (range 1 5)
                     d (range 1 (inc depth))]
                 {[[from d] (char (+ 64 to))] (set
                                                (concat
                                                  (map #(vector from %) (range 1 d))
                                                  (range (inc (min from to)) (inc (max from to)))
                                                  (map #(vector to %) (range 1 (inc depth)))))})))

(defn generate-indirect-blocks
  [to-fn blocks-fn]
  (apply merge (for [from (range 1 5)
                     to (to-fn from)
                     :let [blocks (set (blocks-fn from to))]]
                 (into
                   {[to (char (+ 64 from))] (into (disj blocks to) (map #(vector from %) (range 1 (inc depth))))}
                   (map
                     (fn [d] [[[from d] to] (into blocks (map #(vector from %) (range 1 d)))])
                     (range 1 (inc depth)))))))

(def blocks
  (merge
    direct-blocks
    (generate-indirect-blocks #(range 0 (inc %)) #(range %2 (inc %1)))
    (generate-indirect-blocks #(range (inc %) 7) #(range (inc %1) (inc %2)))))

(def amphipod-types
  (apply vector (for [y (range depth)
                      x (range 4)]
                  (get-in input [(+ 2 y) (+ 3 (* 2 x))]))))

(def initial-positions
  (apply vector (for [y (range depth)
                      x (range 4)]
                  [(inc x) (inc y)])))

(defn compute-cost
  ([cost-per-step from to]
   (* cost-per-step (+
                      (second from)
                      (* 2 (Math/abs (- (first from) to))))))
  ([cost-per-step from to via]
   (* cost-per-step (+
                      (if (contains? #{0 6} via) -2 0)
                      (second from)
                      (if (<= via (first from)) (inc (* 2 (- (first from) via))) (dec (* 2 (- via (first from)))))
                      (if (<= via to) (inc (* 2 (- to via))) (dec (* 2 (- via to))))))))

(defn compute-paths
  [from to-letter]
  (let [to-number (- (int to-letter) 64)
        cost-per-step (reduce (fn [a _] (* 10 a)) 1 (range 1 to-number))
        p1 (range 0 (inc (min (first from) to-number)))
        p2 (range (inc (max (first from) to-number)) 7)]
    (conj (map
            #(vector (compute-cost cost-per-step from to-number %) [from %] [% to-letter])
            (concat p1 p2))
          [(compute-cost cost-per-step from to-number) [from to-letter]])))

(def paths
  (apply vector (map compute-paths initial-positions amphipod-types)))

(defn available-paths
  [positions paths]
  (map #(filter
          (fn [[_ step]] (not-any? (or (blocks step) #{}) positions))
          %)
       paths))

(defn my-min
  [xs]
  (if (empty? xs) Long/MAX_VALUE (apply min xs)))

(defn find-finishable-path
  [aap]
  (some identity
        (map-indexed
          (fn [pod-num ap]
            (if-let [finishable-path (some #(if (-> % second second char?) %) ap)]
              [pod-num finishable-path]))
          aap)))

(defn sort-amphipods
  [positions paths summed-cost]
  (if (every? char? positions)
    summed-cost
    (let [aap (available-paths positions paths)]
      (if-let [[finishable-pod-num finishable-path] (find-finishable-path aap)]
        (recur
          (assoc positions finishable-pod-num (second (second finishable-path)))
          (assoc paths finishable-pod-num ())
          (+ summed-cost (first finishable-path)))
        (my-min
          (apply concat
            (map-indexed
              (fn [pod-num ap]
                (apply vector (map
                                #(sort-amphipods
                                    (assoc positions pod-num (-> % second second))
                                    (assoc paths pod-num [[0 (last %)]])
                                    (+ summed-cost (first %)))
                                ap)))
              aap)))))))

(def base-cost
  (apply + (map (partial * 1111)  (range 1 (inc depth)))))

(println (+ base-cost (sort-amphipods initial-positions paths 0)))
