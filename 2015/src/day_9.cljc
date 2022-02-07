(ns day-9 (:require util))

(def input
  (util/read-input #(re-find (re-matcher #"(\w+) to (\w+) = (\d+)" %))))

(def distances
  (reduce
    (fn [m [_ a b d]] (let [d (Integer/parseInt d)] (-> m (assoc [a b] d) (assoc [b a] d))))
    {}
    input))

(def cities
  (->> distances keys (map first) set))

(defn compute-routes
  [cities]
  (if (empty? cities)
    [()]
    (for [city cities
          route (compute-routes (disj cities city))]
      (cons city route))))

(def routes
  (compute-routes cities))

(defn total-distance
  [route]
  (apply + (map distances (map vector route (next route)))))

(println (apply min (map total-distance routes)))
(println (apply max (map total-distance routes)))
