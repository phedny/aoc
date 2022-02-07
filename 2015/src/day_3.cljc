(ns day-3 (:require util))

(def input
  (first (util/read-input seq)))

(defn move
  [[x y] direction]
  (case direction
    \^ [x (inc y)]
    \v [x (dec y)]
    \> [(inc x) y]
    \< [(dec x) y]))

(defn traverse
  [positions directions]
  (if (empty? directions)
    (list positions)
    (lazy-seq (cons positions (traverse (map move positions (first directions)) (rest directions))))))

(defn compute-result
  [n]
  (->> input (partition n) (traverse (repeat n [0 0])) (mapcat identity) set count))

(println (compute-result 1))
(println (compute-result 2))
