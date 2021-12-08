(ns day-8 (:require util [clojure.set :as set]))

(def input
  (apply util/read-input #"\n" #" " (repeat 15 (comp set seq))))

(defn common-led-count
  [candidates digit c]
  (first (filter #(= c (count (set/intersection digit %))) candidates)))

(defn additional-led-count
  [candidates digit c]
  (first (filter #(= c (count (set/difference % digit))) candidates)))

(defn recognise-digits
  [leds]
  (let [grouped-by-count (group-by count leds)
        digit-1 (first (grouped-by-count 2))
        digit-7 (first (grouped-by-count 3))
        digit-4 (first (grouped-by-count 4))
        digit-8 (first (grouped-by-count 7))
        digit-2 (common-led-count (grouped-by-count 5) digit-4 2)
        digit-3 (common-led-count (grouped-by-count 5) digit-1 2)
        digit-5 (first (disj (set (grouped-by-count 5)) digit-2 digit-3))
        digit-9 (additional-led-count (grouped-by-count 6) digit-3 1)
        digit-6 (additional-led-count (disj (set (grouped-by-count 6)) digit-9) digit-5 1)
        digit-0 (first (disj (set (grouped-by-count 6)) digit-6 digit-9))]
    {digit-0 0 digit-1 1 digit-2 2 digit-3 3 digit-4 4 digit-5 5 digit-6 6 digit-7 7 digit-8 8 digit-9 9}))

(defn line->digits
  [line]
  (map (recognise-digits (take 10 line)) (drop 11 line)))

; Part a
(defn count-trivial
  [line]
  (->> line line->digits (filter #(contains? #{1 4 7 8} %)) count))

(println (apply + (map count-trivial input)))

; Part b
(defn line->number
  [line]
  (->> line line->digits (apply str) Integer/parseInt))

(println (apply + (map line->number input)))
