(ns day-12 (:require util [clojure.data.json :as json]))

(def input
  (first (util/read-input)))

(defn extract-numbers-a
  [d]
  (cond
    (number? d) [d]
    (map? d) (mapcat extract-numbers-a (vals d))
    (vector? d) (mapcat extract-numbers-a d)))

(defn extract-numbers-b
  [d]
  (cond
    (number? d) [d]
    (and (map? d) (not-any? #{"red"} (vals d))) (mapcat extract-numbers-b (vals d))
    (vector? d) (mapcat extract-numbers-b d)))

(println (apply + (extract-numbers-a (json/read-str input))))
(println (apply + (extract-numbers-b (json/read-str input))))
