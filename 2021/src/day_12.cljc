(ns day-12 (:require util))

(def input
  (util/read-input #"\n" #"-" identity identity))

(defn add-to-system
  [system [from to]]
  (update-in system [from] #(if (nil? %) (list to) (cons to %))))

(defn add-to-system-bidirectional
  [system [from to]]
  (add-to-system (add-to-system system [from to]) [to from]))

(def cave-system
  (reduce add-to-system-bidirectional {} input))

(defn is-small-cave
  [cave]
  (Character/isLowerCase (first (seq cave))))

(defn is-large-cave
  [cave]
  (Character/isUpperCase (first (seq cave))))

(defn small-cave-revisit-available
  [prefix]
  (let [small-caves (filter is-small-cave prefix)]
    (= (count small-caves) (count (set small-caves)))))

(defn can-go-to
  [allowed-single-small-revisits prefix step]
  (or
    (is-large-cave step)
    (and
      (small-cave-revisit-available prefix)
      (= (count (filter #(= % step) prefix)) allowed-single-small-revisits)
      (not= "start" step))
    (not (some #(= % step) prefix))))

(defn traverse-cave-system
  [system allowed-single-small-revisits prefix]
  (let [candidates (filter (partial can-go-to allowed-single-small-revisits prefix) (system (first prefix)))]
    (cond
      (= "end" (first prefix)) 1
      (empty? candidates) 0
      :else (apply + (map #(traverse-cave-system system allowed-single-small-revisits (cons % prefix)) candidates)))))

(def result-part-a
  (traverse-cave-system cave-system 0 (list "start")))

(println result-part-a)

(def result-part-b
  (traverse-cave-system cave-system 1 (list "start")))

(println result-part-b)
