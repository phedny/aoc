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

(defn winning-boards
  [numbers boards]
  (filter (partial winning-board? numbers) boards))

(defn losing-boards
  [numbers boards]
  (filter (complement (partial winning-board? numbers)) boards))

(defn compute-answer
  [board drawn-numbers]
  (let [sum-of-unmarked-answers (apply + (set/difference (set (flatten board)) (set drawn-numbers)))]
    (* sum-of-unmarked-answers (last drawn-numbers))))

; Part a
(def shortest-seq-of-prefixes-with-winning-board
  (first (drop-while #(empty? (winning-boards % bingo-boards)) drawn-number-prefixes)))

(def winning-board
  (first (winning-boards shortest-seq-of-prefixes-with-winning-board bingo-boards)))

(println (compute-answer winning-board shortest-seq-of-prefixes-with-winning-board))

; Part b
(def shortest-seq-of-prefixes-with-only-winning-boards
  (first (drop-while #(< (count (winning-boards % bingo-boards)) (count bingo-boards)) drawn-number-prefixes)))

(def losing-board
 (first (losing-boards (drop-last shortest-seq-of-prefixes-with-only-winning-boards) bingo-boards)))

(println (compute-answer losing-board shortest-seq-of-prefixes-with-only-winning-boards))
