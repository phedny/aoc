(ns day-5 (:require util [clojure.string :as str]))

(defn str->coord
  [s]
  (apply vector (map #(Integer/parseInt %) (str/split s #","))))

(def input
  (util/read-input #"\n" #" " str->coord identity str->coord))

(defn auto-range
  [start end]
  (cond
    (< start end) (range start (inc end))
    (> start end) (range start (dec end) -1)
    :else (repeat start)))

(defn draw-line
  [[[x1 y1] _ [x2 y2]]]
  (map vector (auto-range x1 x2) (auto-range y1 y2)))

(defn straight-line?
  [[[x1 y1] _ [x2 y2]]]
  (or (= x1 x2) (= y1 y2)))

(defn compute-answer
  [coords]
  (->> coords (mapcat draw-line) (group-by identity) vals (map count) (filter #(> % 1)) count))

(println (compute-answer (filter straight-line? input)))
(println (compute-answer input))
