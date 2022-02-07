(ns day-2 (:require util))

(def input
  (apply util/read-input #"\n" #"x" (repeat 3 #(Integer/parseInt %))))

(defn wrapping-paper-area
  [[l w h]]
  (let [areas [(* l w) (* l h) (* w h)]]
    (apply + (apply min areas) (map (partial * 2) areas))))

(defn ribbon-length
  [lengths]
  (apply + (apply * lengths) (map (partial * 2) (take 2 (sort lengths)))))

(defn compute-result
  [f]
  (apply + (map f input)))

(println (compute-result wrapping-paper-area))
(println (compute-result ribbon-length))
