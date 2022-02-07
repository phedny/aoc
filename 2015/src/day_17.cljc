(ns day-17 (:require util))

(def input
  (util/read-input #(Long/parseLong %)))

(defn powerset
  [[x & xs]]
  (if x
    (mapcat #(vector % (conj % x)) (powerset xs))
    [()]))

(def valid-combinations
  (filter #(= (apply + %) 150) (powerset input)))

(println (count valid-combinations))
(println (let [cs (map count valid-combinations)
               min-c (apply min cs)]
           (count (filter #{min-c} cs))))
