(ns day-1 (:require util))

(def input
  (first (util/read-input seq)))

(def result-part-a
  (- (count (filter #(= % \() input)) (count (filter #(= % \)) input))))

(println result-part-a)

(defn to-basement
  [n i]
  (if (< n 0)
    0
    (inc (to-basement ((if (= (first i) \() inc dec) n) (rest i)))))

(def result-part-b
  (to-basement 0 input))

(println result-part-b)
