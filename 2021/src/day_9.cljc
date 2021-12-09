(ns day-9 (:require util [clojure.set :as set]))

(def input
  (apply vector (util/read-input (comp (partial apply vector) (partial map #(Character/digit % 10)) seq))))

(def height
  (count input))

(def width
  (count (first input)))

(defn neighbors
  [row col]
  (filter (complement #(nil? (get-in input %))) [[(dec row) col] [(inc row) col] [row (dec col)] [row (inc col)]]))

(defn low-point?
  [row col]
  (let [val (get-in input [row col])]
    (every? #(< val %) (map (partial get-in input) (neighbors row col)))))

(def low-points
  (for [row (range 0 height)
        col (range 0 width)
        :when (low-point? row col)]
    [row col]))

(def result-part-a
  (apply + (map #(inc (get-in input %)) low-points)))

(println result-part-a)

(defn low-point->basin
  [row col]
  {:area #{[row col]} :boundary #{[row col]}})

(defn higher-neighbors
  [row col]
  (let [val (get-in input [row col])]
    (filter (every-pred #(< val (get-in input %)) #(not= 9 (get-in input %))) (neighbors row col))))

(defn extend-basin
  [{area :area boundary :boundary}]
  (let [new-boundary (set/difference (set (mapcat #(apply higher-neighbors %) boundary)) area)]
    {:area (set/union area new-boundary) :boundary new-boundary}))

(defn find-basin*
  [{area :area boundary :boundary :as basin}]
  (if (empty? boundary)
    (list)
    (->> basin extend-basin find-basin* (cons area) lazy-seq)))

(defn find-basin
  [basin]
  (last (find-basin* basin)))

(def basin-sizes
  (map (comp count find-basin (partial apply low-point->basin)) low-points))

(def result-part-b
  (apply * (take 3 (sort > basin-sizes))))

(println result-part-b)
