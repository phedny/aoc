(ns day-24 (:require util [clojure.set :as set]))

(def input (sort > (util/read-input #(Integer/parseInt %))))

(defn find-groups
  [goal available [x & xs]]
  (cond
    (zero? goal) [#{}]
    (neg? goal) []
    (nil? x) []
    (> goal available) []
    (= goal available) [(conj (set xs) x)]
    (< goal x) (recur goal (- available x) xs)
    :else (concat (map #(conj % x) (find-groups (- goal x) (- available x) xs)) (find-groups goal (- available x) xs))))

(defn find-balanced
  [num-groups groups balanced-groups]
  (if (= num-groups (count balanced-groups))
    [(reverse balanced-groups)]
    (let [candidates (->> groups
                          (filter #(not= (first balanced-groups) %))
                          (filter #(empty? (set/intersection (first balanced-groups) %))))]
      (mapcat #(find-balanced num-groups candidates (cons % balanced-groups)) candidates))))

(defn find-best-balance
  [num-groups]
  (let [groups (find-groups (/ (reduce + input) num-groups) (reduce + input) input)
        group-1-count (apply min (map count groups))
        group-1s (sort-by #(reduce * %) (filter #(= group-1-count (count %)) groups))]
    (first (first (filter not-empty (map #(find-balanced num-groups groups (list %)) group-1s))))))

(println (reduce * (first (find-best-balance 3))))
(println (reduce * (first (find-best-balance 4))))
