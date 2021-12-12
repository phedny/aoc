(ns day-12 (:require util))

(defn add-to-system
  [system [from to]]
  (let [system-with-cave (cond
                           (contains? system from) system
                           (= "start" from) (assoc system "start" {:start true :to (list)})
                           (= "end" from) (assoc system "end" {:end true :to (list)})
                           (Character/isUpperCase (first (seq from))) (assoc system from {:large true :to (list)})
                           :else (assoc system from {:small true :to (list)}))]
    (update-in system-with-cave [from :to] (partial cons to))))

(defn add-to-system-bidirectional
  [system [from to]]
  (-> system (add-to-system [from to]) (add-to-system [to from])))

(def system
  (reduce
    add-to-system-bidirectional
    {}
    (util/read-input #"\n" #"-" identity identity)))

(defn small-cave-revisit-available
  [prefix]
  (let [small-caves (filter #((system %) :small) prefix)]
    (= (count small-caves) (count (set small-caves)))))

(defn large-cave?
  [[_ cave]]
  ((system cave) :large))

(defn not-visited-before?
  [[prefix cave]]
  (not (some #(= % cave) prefix)))

(defn small-cave-first-revisit?
  [[prefix cave]]
  (and
    (not ((system cave) :start))
    (small-cave-revisit-available prefix)
    (= (count (filter #(= % cave) prefix)) 1)))

(defn traverse-cave-system
  [continue-to? prefix]
  (let [candidates (filter (partial (comp continue-to? vector) prefix) (:to (system (first prefix))))]
    (cond
      ((system (first prefix)) :end) 1
      (empty? candidates) 0
      :else (apply + (map #(traverse-cave-system continue-to? (cons % prefix)) candidates)))))

(println
  (traverse-cave-system (some-fn large-cave? not-visited-before?) (list "start")))

(println
  (traverse-cave-system (some-fn large-cave? not-visited-before? small-cave-first-revisit?) (list "start")))
