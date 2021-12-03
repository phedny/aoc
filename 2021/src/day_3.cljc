(ns day-3 (:require util))

(def chars->ints
  (partial map #(Character/digit % 10)))

(def input
  (util/read-input (comp chars->ints seq)))

(defn transpose
  [xs]
  (apply map vector xs))

(defn count-bit-values
  [values]
  (reduce
    #(update-in %1 [%2] inc)
    [0 0]
    values))

(defn compare-bits
  [more-zeroes more-ones bit-string]
  (if (apply > (count-bit-values bit-string)) more-zeroes more-ones))

(def most-common-bit
  (partial compare-bits 0 1))

(def least-common-bit
  (partial compare-bits 1 0))

(defn parse-binary
  [bits]
  (Integer/parseInt (apply str bits) 2))

(defn compute-answer
  [fn bit-strings]
  (* (fn most-common-bit bit-strings) (fn least-common-bit bit-strings)))

; Part a
(defn merge-bit-strings
  [fn bit-strings]
  (parse-binary (map fn (transpose bit-strings))))

(println (compute-answer merge-bit-strings input))

; Part b

(defn apply-bit-criteria
  ([fn values]
   (parse-binary (apply-bit-criteria fn values 0)))
  ([fn values bit]
   (if (= (count values) 1)
     (first values)
     (let [v (->> values transpose (map fn))]
       (recur fn (filter #(= (nth % bit) (nth v bit)) values) (inc bit))))))

(println (compute-answer apply-bit-criteria input))
