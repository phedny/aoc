(ns day-13 (:require util))

(def input
  (util/read-input))

(defn parse-dot
  [line]
  (apply vector (map #(Integer/parseInt %) (clojure.string/split line #","))))

(def dots
  (set (map parse-dot (take-while (complement empty?) input))))

(defn parse-fold
  [line]
  (let [groups (re-find (re-matcher #"fold along ([xy])=(\d+)" line))]
    [(first (get groups 1)) (Integer/parseInt (get groups 2))]))

(def folds
  (map parse-fold (rest (drop-while (complement empty?) input))))

(defn fold
  [dots [fold-direction fold-line]]
  (let [fold-index (- (int fold-direction) 120)
        fold-fn #(- fold-line (Math/abs (- fold-line %)))]
    (set (map #(update-in % [fold-index] fold-fn) dots))))

(def result-part-a
  (->> folds first (fold dots) count))

(println result-part-a)

(defn print-dots
  [dots]
  (let [max-x (apply max (map first dots))
        max-y (apply max (map second dots))]
    (doseq [y (range (inc max-y))]
      (println (apply str (map #(if (contains? dots [% y]) \# \ ) (range (inc max-x))))))))

(def folded-manual
  (reduce fold dots folds))

(print-dots folded-manual)
