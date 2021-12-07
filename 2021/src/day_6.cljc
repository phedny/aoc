(ns day-6 (:require util))

(def input
  (util/read-input #"," #(Long/parseLong %)))

(def initial-fish-timers
  (map (fn [timer] (count (filter #(= % timer) input))) (range 0 9)))

(defn update-timers
  [timers]
  (let [spawn-count (first timers)
        rotated-timers (apply vector (concat (rest timers) [spawn-count]))]
    (update-in rotated-timers [6] + spawn-count)))

(defn fish-count*
  [timers]
  (lazy-seq (cons (apply + timers) (fish-count* (update-timers timers)))))

(def fish-count
  (fish-count* initial-fish-timers))

(println (nth fish-count 80))
(println (nth fish-count 256))
