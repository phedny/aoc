(ns day-16 (:require util))

(def input
  (first (util/read-input)))

(defn hex->bits
  [h]
  (let [b (map #(- (byte %) 48) (Integer/toString (Character/digit h 16) 2))]
    (concat (repeat (- 4 (count b)) 0) (seq b))))

(defn string->byte-array
  [s]
  (byte-array (* 4 (count s)) (mapcat hex->bits (seq s))))

(defn string->byte-buffer
  [s]
  (java.nio.ByteBuffer/wrap (string->byte-array s)))

(defn read-number
  [bits bb]
  (reduce (fn [acc _] (+ (* acc 2) (.get bb))) 0 (range bits)))

(defn read-literal
  ([bb] (read-literal bb 0))
  ([bb n]
   (let [more? (= 1 (.get bb))
         value (+ (* n 16) (read-number 4 bb))]
     (if more? (recur bb value) value))))

(declare byte-buffer->packets)
(defn read-operator
  [bb]
  (if (= 0 (.get bb))
    (let [length (read-number 15 bb)
          sub-bb (doto (.slice bb) (.limit length))]
      (.position bb (+ (.position bb) length))
      (byte-buffer->packets sub-bb))
    (let [count (read-number 11 bb)]
      (byte-buffer->packets bb count))))

(defn byte-buffer->packet
  [bb]
  (let [version (read-number 3 bb)
        type-id (read-number 3 bb)
        packet {:version version :type-id type-id}]
    (if (= 4 type-id)
      (assoc packet :value (read-literal bb))
      (assoc packet :sub-packets (read-operator bb)))))

(defn byte-buffer->packets
  ([bb]
   (if (.hasRemaining bb)
     (cons (byte-buffer->packet bb) (byte-buffer->packets bb))
     ()))
  ([bb n]
   (if (> n 0)
     (cons (byte-buffer->packet bb) (byte-buffer->packets bb (dec n)))
     ())))

(defn version-sum
  [{version :version sub-packets :sub-packets}]
  (apply + version (map version-sum (or sub-packets ()))))

(println (-> input string->byte-buffer byte-buffer->packet version-sum))

(defn packet->clojure
  [packet]
  (case (packet :type-id)
    0 (apply list '+ (map packet->clojure (packet :sub-packets)))
    1 (apply list '* (map packet->clojure (packet :sub-packets)))
    2 (apply list 'min (map packet->clojure (packet :sub-packets)))
    3 (apply list 'max (map packet->clojure (packet :sub-packets)))
    4 (packet :value)
    5 (list 'if (apply list '> (map packet->clojure (packet :sub-packets))) 1 0)
    6 (list 'if (apply list '< (map packet->clojure (packet :sub-packets))) 1 0)
    7 (list 'if (apply list '= (map packet->clojure (packet :sub-packets))) 1 0)))

;(println (-> input string->byte-buffer byte-buffer->packet packet->clojure))
(println (-> input string->byte-buffer byte-buffer->packet packet->clojure eval))
