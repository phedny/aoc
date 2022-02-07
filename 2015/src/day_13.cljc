(ns day-13 (:require util))

(def input
  (reduce
    (fn [m [_ subj gl amt obj]] (assoc-in m [subj obj] (* (if (= "gain" gl) 1 -1) (Long/parseLong amt))))
    {}
    (util/read-input #(first (re-seq #"(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)." %)))))

(defn permutations
  [xs]
  (if (next xs)
    (mapcat #(map (partial cons (nth xs %)) (permutations (concat (take % xs) (drop (inc %) xs)))) (range (count xs)))
    [xs]))

(defn happiness-change-of-seating-arrangement
  [names]
  (apply + (map #(+ (get-in input [%1 %2] 0) (get-in input [%2 %1] 0)) names (concat (rest names) [(first names)]))))

(defn max-happiness-change-for
  [names]
  (apply max (map happiness-change-of-seating-arrangement (permutations names))))

(println (max-happiness-change-for (keys input)))
(println (max-happiness-change-for (conj (keys input) "")))
