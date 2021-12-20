(ns day-20 (:require util))

(def input
  (util/read-input))

(defn char->int
  [c]
  (if (= c \#) 1 0))

(def mapping
  (->> input first seq (map char->int) (apply vector)))

(def initial-pattern
  [0 (->> input (drop 2) (map #(->> % seq (map char->int))))])

(defn extend-pattern
  [inf pattern]
  (let [with-empty-cols (map #(concat [inf inf] % [inf inf]) pattern)
        empty-rows (repeat 2 (repeat (count (first with-empty-cols)) inf))]
    (concat empty-rows with-empty-cols empty-rows)))

(defn substitute
  [args]
  (get mapping (Integer/parseInt (apply str (flatten args)) 2)))

(defn substitute-row
  [row]
  (map substitute row))

(defn substitute-pattern
  [[inf pattern]]
  (let [partitioned-pattern (partition 3 1 (map (partial partition 3 1) (extend-pattern inf pattern)))
        transposed (map #(apply map vector %) partitioned-pattern)
        substituted-inf (get mapping (if (= 0 inf) 0 511))]
    [substituted-inf (map substitute-row transposed)]))

(def result-part-a
  (->> initial-pattern substitute-pattern substitute-pattern flatten (filter #(= % 1)) count))

(println result-part-a)
