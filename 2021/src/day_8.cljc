(ns day-8 (:require util [clojure.set :as set]))

(def input
  (apply util/read-input #"\n" #" " (repeat 15 (comp set seq))))

; Part a
(def inputs-by-led-count
  (->> input (mapcat identity) (group-by count)))

(defn count-digits-with-led-count
  [led-count]
  (- (count (inputs-by-led-count led-count)) (count input)))

(def result-part-a
  (apply + (map count-digits-with-led-count [2 3 4 7])))

(println result-part-a)

; Part b
(defn differentiate-2-3-5
  [candidates digit-1 digit-4]
  (let [the-2 (first (filter #(= 2 (count (set/intersection % digit-4))) candidates))
        the-3 (first (filter #(= 2 (count (set/intersection % digit-1))) candidates))]
    {2 the-2 3 the-3 5 (first (disj (set candidates) the-2 the-3))}))

(defn find-one-larger
  [candidates digit]
  (first (filter #(= 1 (count (set/difference % digit))) candidates)))

(defn differentiate-0-6-9
  [candidates digit-3 digit-5]
  (let [the-9 (find-one-larger candidates digit-3)
        the-6 (find-one-larger (disj (set candidates) the-9) digit-5)]
    {6 the-6 9 the-9 0 (first (disj (set candidates) the-6 the-9))}))

(defn map-digits
  [leds]
  (let [grouped-by-count (group-by count leds)
        trivial {1 (first (grouped-by-count 2)) 4 (first (grouped-by-count 4)) 7 (first (grouped-by-count 3)) 8 (first (grouped-by-count 7))}
        the-2-3-5 (differentiate-2-3-5 (grouped-by-count 5) (trivial 1) (trivial 4))
        the-0-6-9 (differentiate-0-6-9 (grouped-by-count 6) (the-2-3-5 3) (the-2-3-5 5))]
    (clojure.set/map-invert (merge trivial the-2-3-5 the-0-6-9))))

(defn line->number
  [line]
  (let [digits->number (map-digits (take 10 line))]
    (Integer/parseInt (apply str (map digits->number (drop 11 line))))))

(def result-part-b
  (apply + (map line->number input)))

(println result-part-b)
