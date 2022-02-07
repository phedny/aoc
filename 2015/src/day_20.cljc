(ns day-20)

(declare divisors)
(defn divisors*
  [n]
  (if-let [sd (some #(if (zero? (mod n %)) %) (range 2 (Math/sqrt (inc n))))]
    (conj (set (mapcat (fn [x] [x (* x sd)]) (divisors (quot n sd)))) n sd)
    #{1 n}))
(def divisors (memoize divisors*))

(defn find-the-house
  [f]
  (some #(if (>= (f %) 34000000) %) (iterate inc 2)))

(println (find-the-house (fn [house] (* 10 (reduce + (divisors house))))))
(println (find-the-house (fn [house] (* 11 (reduce + (filter #(<= house (* % 50)) (divisors house)))))))
