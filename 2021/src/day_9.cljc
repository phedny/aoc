(ns day-9 (:require util))

(def input
  (apply vector (util/read-input (comp (partial apply vector) (partial map #(Character/digit % 10)) seq))))

(def height
  (count input))

(def width
  (count (first input)))

(defn neighbors
  [row col]
  (map (partial get-in input) [[(dec row) col] [(inc row) col] [row (dec col)] [row (inc col)]]))

(defn low-point?
  [row col]
  (let [val (get-in input [row col])]
    (every? #(< val %) (filter (complement nil?) (neighbors row col)))))

(def low-points
  (for [row (range 0 height)
        col (range 0 width)
        :when (low-point? row col)]
    [row col]))

(def result-part-a
  (apply + (map #(inc (get-in input %)) low-points)))

(println result-part-a)