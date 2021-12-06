(ns day-6 (:require util))

(def input
  (->> (util/read-input #(clojure.string/split % #",")) first (map #(Integer/parseInt %))))

(defn spawn-fish-count
  ([]
   (spawn-fish-count (map (fn [n] (count (filter #(= % n) input))) (range 7 -2 -1))))
  ([just-before]
   (lazy-seq (cons (nth just-before 7) (spawn-fish-count (cons (+ (nth just-before 6) (nth just-before 8)) just-before))))))

(defn fish-count
  ([]
   (fish-count (count input) (spawn-fish-count)))
  ([c spawned-fish]
   (lazy-seq (cons c (fish-count (+ c (first spawned-fish)) (rest spawned-fish))))))

(println (nth (fish-count) 80))
(println (nth (fish-count) 256))
