(ns day-7 (:require util))

(def input
  (util/read-input #"," #(Integer/parseInt %)))

(defn cost-of
  [cost-fn positions target]
  (apply + (map #(cost-fn (Math/abs (- target %))) positions)))

(defn lowest-cost
  [cost-fn positions]
  (let [range (range (apply min positions) (inc (apply max positions)))]
    (apply min (map (partial cost-of cost-fn positions) range))))

(println (lowest-cost identity input))
(println (lowest-cost #(quot (* % (inc %)) 2) input))
