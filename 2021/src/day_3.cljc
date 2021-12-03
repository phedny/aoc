(ns day-3 (:require util))

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

(defn least-common-bit
  [[number-of-zeroes number-of-ones]]
  (if (> number-of-zeroes number-of-ones) \1 \0))

; Part a
(def gamma-rate
  (apply str (map (comp most-common-bit count-bit-values) (transpose input))))

(def epsilon-rate
  (apply str (map (comp least-common-bit count-bit-values) (transpose input))))

(def answer-a
  (* (Integer/parseInt gamma-rate 2) (Integer/parseInt epsilon-rate 2)))

(println answer-a)

; Part b

(defn apply-bit-criteria
  [fn values bit]
  (if (= (count values) 1)
    (first values)
    (let [v (->> values transpose (map (comp fn count-bit-values)))]
      (recur fn (filter #(= (nth % bit) (nth v bit)) values) (inc bit)))))

(def oxygen-generator-rating
  (apply str (apply-bit-criteria most-common-bit input 0)))

(def CO2-scrubber-rating
  (apply str (apply-bit-criteria least-common-bit input 0)))

(def answer-b
  (* (Integer/parseInt oxygen-generator-rating 2) (Integer/parseInt CO2-scrubber-rating 2)))

(println answer-b)
