(ns day-3 (:require util))
;(def filename "example")

(def input
  (util/read-input seq))

(defn transpose
  [xs]
  (apply map vector xs))

(defn count-bit-values
  [values]
  (reduce
    (fn
      [counts value]
      (update-in counts [(Character/digit value 10)] inc))
    [0 0]
    values))

(defn most-common-bit
  [[number-of-zeroes number-of-ones]]
  (if (> number-of-zeroes number-of-ones) \0 \1))

(defn least
  [[number-of-zeroes number-of-ones]]
  (if (< number-of-zeroes number-of-ones) \0 \1))

(def gamma-rate
  (apply str (map (comp most-common-bit count-bit-values) (transpose input))))

(def epsilon-rate
  (apply str (map (comp least count-bit-values) (transpose input))))

(def answer
  (* (Integer/parseInt gamma-rate 2) (Integer/parseInt epsilon-rate 2)))

(println answer)
