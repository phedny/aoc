(ns day-1 (:require util))

(def input
  (util/read-input #(Integer/parseInt %)))

(defn sliding-sums
  [n xs]
  (map #(apply + %) (partition n 1 xs)))

(defn count-increasing
  [xs]
  (count (filter identity (map < xs (rest xs)))))

(println (count-increasing input))
(println (count-increasing (sliding-sums 3 input)))
