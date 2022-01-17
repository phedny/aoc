(ns day-14 (:require util))

(def input
  (map
    (comp (partial map #(Long/parseLong %)) rest)
    (util/read-input #(first (re-seq #"\w+ can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds." %)))))

(defn advance-amount
  [second [run-speed run-time rest-time]]
  (if (< (mod second (+ run-time rest-time)) run-time) run-speed 0))

(defn position*
  ([reindeer]
   (position* reindeer 0 0))
  ([reindeer second position]
   (lazy-seq (cons position (position* reindeer (inc second) (+ position (advance-amount second reindeer)))))))

(def positions
  (map position* input))

(defn points*
  ([positions]
   (points* positions (map (constantly 0) positions)))
  ([positions points]
   (lazy-seq (let [current-positions (map second positions)
                   max-position (apply max current-positions)]
               (cons points (points* (map rest positions) (map #(+ (if (= %1 max-position) 1 0) %2) current-positions points)))))))

(println (apply max (map #(nth % 2503) positions)))
(println (apply max (nth (points* positions) 2503)))
