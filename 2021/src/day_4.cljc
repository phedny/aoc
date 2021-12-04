(ns day-4 (:require util [clojure.string :as str] [clojure.set :as set]))

(def input
  (util/read-input))

(def drawn-numbers
  (as-> input $ (first $) (str/split $ #",") (map #(Integer/parseInt %) $)))

; Parse the bingo boards
(defn parse-bingo-line
  [line]
  (as-> line $ (str/trim $) (str/split $ #" +") (map #(Integer/parseInt %) $)))

(def bingo-boards
  (as-> input $ (rest $) (partition 6 $) (map #(map parse-bingo-line (rest %)) $)))

; Play the game
(defn transpose
  [xs]
  (apply map vector xs))

(defn winning-line?
  [numbers line]
  (every? (fn [n] (some #(= % n) numbers)) line))

(defn winning-board?
  [numbers board]
  (some (partial winning-line? numbers) (concat board (transpose board))))

(def drawn-number-prefixes
  (map #(take % drawn-numbers) (rest (range))))

(defn pick-winning-board
  [numbers boards]
  (some (partial winning-board? numbers) boards))

(def shortest-seq-of-prefixes-with-winning-board
  (first (drop-while (complement #(pick-winning-board % bingo-boards)) drawn-number-prefixes)))

(def winning-board
  (first (filter #(winning-board? shortest-seq-of-prefixes-with-winning-board %) bingo-boards)))

(def unmarked-numbers
  (set/difference (set (flatten winning-board)) (set shortest-seq-of-prefixes-with-winning-board)))

(def result
  (* (apply + unmarked-numbers) (last shortest-seq-of-prefixes-with-winning-board)))

(println result)
