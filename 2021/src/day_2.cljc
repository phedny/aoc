(ns day-2 (:require util))

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
        "forward" (-> position
                    (update-in [:x] + distance)
                    (update-in [:y] + (* (position :aim) distance)))
        "down" (update-in position [:aim] #(+ % distance))
        "up" (update-in position [:aim] #(- % distance))))
    {:x 0 :y 0 :aim 0}
    input))

(println (* (answer :x) (answer :y)))
