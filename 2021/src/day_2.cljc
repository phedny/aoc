(ns day-2 (:require util))

(def input
  (util/read-input #"\n" #" " identity #(Integer/parseInt %)))

; Instruction implementation for part a
(defn forward-a
  [position distance]
  (update-in position [:x] + distance))

(defn down-a
  [position distance]
  (update-in position [:y] + distance))

(defn up-a
  [position distance]
  (update-in position [:y] #(- % distance)))

; Instruction implementation for part b
(defn forward-b
  [position distance]
  (-> position
      (update-in [:x] + distance)
      (update-in [:y] + (* (position :aim) distance))))

(defn down-b
  [position distance]
  (update-in position [:aim] + distance))

(defn up-b
  [position distance]
  (update-in position [:aim] #(- % distance)))

; Generic code to iterate over instructions
(defn execute-instructions
  [part]
  (reduce
    (fn
      [position [direction distance]]
      ((->> (str direction "-" part) symbol resolve) position distance))
    {:x 0 :y 0 :aim 0}
    input))

; Print the results
(def answer-a (execute-instructions "a"))
(println (* (answer-a :x) (answer-a :y)))

(def answer-b (execute-instructions "b"))
(println (* (answer-b :x) (answer-b :y)))
