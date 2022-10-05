(ns day-3 (:require util))

(def input (util/read-input #"\n" #" +" identity parse-long parse-long parse-long))

(defn valid-triangle?
  [a b c]
  (pos? (- (+ a b c) (* 2 (max a b c)))))

(defn count-valid-triangles
  [input]
  (count (filter #(apply valid-triangle? %) input)))

(println (count-valid-triangles (map rest input)))

(defn transpose
  [xs]
  (apply map vector xs))

(println (count-valid-triangles (mapcat transpose (partition 3 (map rest input)))))
