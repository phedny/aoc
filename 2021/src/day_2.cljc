(ns day-2 (:require util))
;(def filename "example")

(def input
  (util/read-input
    (fn
      [line]
      (update-in (clojure.string/split line #" ") [1] #(Integer/parseInt %)))))

(def answer
  (reduce
    (fn
      [position [direction distance]]
      (case direction
        "forward" (update-in position [:x] #(+ % distance))
        "down" (update-in position [:y] #(+ % distance))
        "up" (update-in position [:y] #(- % distance))))
    {:x 0 :y 0}
    input))

(println (* (answer :x) (answer :y)))
