(ns day-7 (:require util))

(def input
  (first (util/read-input (fn [line] (map #(Integer/parseInt %) (clojure.string/split line #","))))))

(defn linear-price
  [to from]
  (Math/abs (- from to)))

(defn triangular-price
  [to from]
  (let [steps (linear-price from to)]
    (quot (* steps (inc steps)) 2)))

(defn cost-of
  [cost-fn positions target]
  (apply + (map (partial cost-fn target) positions)))

(defn lowest-cost
  [cost-fn positions]
  (let [range (range (apply min positions) (inc (apply max positions)))]
    (apply min (map (partial cost-of cost-fn positions) range))))

(println (lowest-cost linear-price input))
(println (lowest-cost triangular-price input))
