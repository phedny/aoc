(ns day-13 (:require [clojure.set :as set]))

(def favorite-number 1352)
(def goal [31 39])

(defn open?
  [[x y]]
  (loop [c 0 n (+ (* x x) (* 3 x) (* 2 x y) y (* y y) favorite-number)]
    (if (zero? n)
      (even? c)
      (recur (if (even? n) c (inc c)) (bit-shift-right n 1)))))

(defn find-shortest-path
  [start]
  (loop [todo (sorted-set [(apply + goal) 0 start])
         seen {}
         shortest-path Integer/MAX_VALUE]
    (let [[estimate cost [x y :as position] :as item] (first todo)
          todo (disj todo item)]
      (cond
        (<= shortest-path estimate) shortest-path
        (= goal position) (recur todo seen cost)
        :else (recur (->> [[x (dec y)] [x (inc y)] [(dec x) y] [(inc x) y]]
                          (filter (fn [[x y]] (and (<= 0 x) (<= 0 y))))
                          (filter open?)
                          (remove #(<= (get seen % Integer/MAX_VALUE) cost))
                          (map #(vector (+ 1 cost (abs (- (first %) (first goal))) (abs (- (second %) (second goal)))) (inc cost) %))
                          (reduce conj todo))
                     (assoc seen position cost)
                     shortest-path)))))

(println (find-shortest-path [1 1]))

(defn count-reachable-cells
  [start max-cost]
  (loop [front #{start}
         seen #{}
         n max-cost]
    (if (zero? n)
      (+ (count front) (count seen))
      (recur (set (for [from front
                        direction [[0 1] [0 -1] [1 0] [-1 0]]
                        :let [x (+ (first from) (first direction)) y (+ (second from) (second direction))]
                        :when (and (<= 0 x) (<= 0 y) (open? [x y]) (not (contains? seen [x y])))]
                    [x y]))
             (set/union seen front)
             (dec n)))))

(println (count-reachable-cells [1 1] 50))
