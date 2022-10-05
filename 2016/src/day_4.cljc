(ns day-4 (:require util))

(def input (util/read-input #(re-matches #"([a-z-]+)-(\d+)\[([a-z]{5})\]" %1)))

(defn valid-pair?
  [freqs [a b]]
  (let [f-a (freqs a 0)
        f-b (freqs b 0)]
    (or (> f-a f-b)
        (and (= f-a f-b) (< (int a) (int b))))))

(defn real-room?
  [[_ name _ csum]]
  (let [freqs (dissoc (frequencies name) \-)]
    (and (every? (partial valid-pair? freqs) (partition 2 1 csum))
         (every? #(valid-pair? freqs [(get csum 4) %]) (keys (apply dissoc freqs csum))))))

(defn decrypt-char
  [c key]
  (if (= c \-)
    \space
    (-> c int (- 97) (+ key) (mod 26) (+ 97) char)))

(defn decrypt
  [[_ name sector]]
  (apply str (map #(decrypt-char % (parse-long sector)) name)))

(println (reduce + (map #(parse-long (get % 2)) (filter real-room? input))))

(println (get (first (filter #(= (decrypt %) "northpole object storage") (filter real-room? input))) 2))