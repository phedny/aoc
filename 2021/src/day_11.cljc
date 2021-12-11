(ns day-11 (:require util [clojure.set :as set]))

(def input
  (apply vector (util/read-input (comp (partial apply vector) (partial map #(Character/digit % 10)) seq))))

(defn neighbors
  [row col]
  (let [coords (for [delta-row (range -1 2)
                     delta-col (range -1 2)
                     :when (or (not= delta-row 0) (not= delta-col 0))]
                 [(+ row delta-row) (+ col delta-col)])]
    (filter (complement #(nil? (get-in input %))) coords)))

(defn coords-that
  [fn grid]
  (for [row (range 0 10)
        col (range 0 10)
        :when (fn (get-in grid [row col]))]
    [row col]))

(defn for-all
  [fn grid]
  (apply vector (map #(apply vector (map fn %)) grid)))

(defn inc-around-flashes
  ([grid] (inc-around-flashes grid (set (coords-that #(> % 9) grid))))
  ([grid flashes] (if (empty? flashes)
                    grid
                    (let [affected-cells (filter
                                           #(< (get-in grid %) 10)
                                           (mapcat #(apply neighbors %) flashes))
                          new-grid (reduce #(update-in %1 %2 inc) grid affected-cells)]
                      (recur new-grid (set/difference (set (coords-that #(> % 9) new-grid)) (set (coords-that #(> % 9) grid))))))))

(defn unflash
  [n]
  (if (> n 9) 0 n))

(def step (comp (partial for-all unflash) inc-around-flashes (partial for-all inc)))

(defn flash-count*
  [grid]
  (let [new-grid (step grid)]
    (lazy-seq (cons (count (coords-that #(= % 0) new-grid)) (flash-count* new-grid)))))

(def flash-count
  (flash-count* input))

(def result-part-a
  (apply + (take 100 flash-count)))

(println result-part-a)

(def result-part-b
  (inc (count (take-while #(< % 100) flash-count))))

(println result-part-b)
