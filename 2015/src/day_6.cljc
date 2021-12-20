(ns day-6 (:require util))

(defn parse-instruction
  [line]
  (let [groups (re-find (re-matcher #"(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)" line))]
    (apply vector (second groups) (map #(Integer/parseInt %) (drop 2 groups)))))

(def input
  (util/read-input parse-instruction))

(defn in-range
  [[x-from y-from x-to y-to] x y]
  (and (<= x-from x) (<= y-from y) (>= x-to x) (>= y-to y)))

(defn instructions->function
  [next instruction->function instructions]
  (if (empty? instructions)
    next
    (let [instruction (first instructions)
          function (instruction->function (first instruction) next)]
      (recur #((if (in-range (rest instruction) %1 %2) function next) %1 %2) instruction->function (rest instructions)))))

(defn compute-result
  [instruction->function]
  (let [light-fn (instructions->function (constantly 0) instruction->function input)
        lights (for [x (range 0 1000) y (range 0 1000)] (light-fn x y))]
    (apply + lights)))

(defn part-a
  [command next]
  (case command
    "turn on" (constantly 1)
    "turn off" (constantly 0)
    "toggle" (comp (partial - 1) next)))

(println (compute-result part-a))

(defn part-b
  [command next]
  (case command
    "turn on" (comp inc next)
    "turn off" (comp #(if (zero? %) 0 (dec %)) next)
    "toggle" (comp (partial + 2) next)))

(println (compute-result part-b))
