(ns day-2 (:require util))

(def input (util/read-input (comp (partial map (comp keyword str)) seq)))

(def moves
  {:U [-1 0] :D [1 0] :L [0 -1] :R [0 1]})

(defn move
  [pad p d]
  (let [np (map + p (moves d))]
    (if (nil? (get-in pad np)) p np)))

(defn compute-code
  [pad start input]
  (->> input
       (reductions (partial reduce (partial move pad)) start)
       next
       (map (partial get-in pad))
       (apply str)))

(def pad-a [[1 2 3]
            [4 5 6]
            [7 8 9]])
(println (compute-code pad-a [1 1] input))

(def pad-b [[nil nil 1 nil nil]
            [nil 2 3 4 nil]
            [5 6 7 8 9]
            [nil \A \B \C nil]
            [nil nil \D nil nil]])
(println (compute-code pad-b [2 0] input))
