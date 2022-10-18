(ns day-10 (:require util))

(def input (util/read-input #(re-matches #"(?:value (\d+) goes to bot (\d+))|(?:bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+))" %)))

(def init
  (reduce #(update %1 (parse-long (get %2 2)) (partial cons (parse-long (get %2 1))))
          {}
          (filter #(get % 1) input)))

(def instrs
  (reduce #(assoc %1 (parse-long (get %2 3)) [(keyword (get %2 4)) (parse-long (get %2 5)) (keyword (get %2 6)) (parse-long (get %2 7))])
          {}
          (filter #(get % 3) input)))

(defn process
  [state [bot-no chips]]
  (-> state
      (update-in (subvec (get instrs bot-no) 0 2) conj (apply min chips))
      (update-in (subvec (get instrs bot-no) 2 4) conj (apply max chips))))

(defn step
  [{:keys [bot output]}]
  (let [to-do (filter #(= 2 (count (second %))) bot)
        state (remove #(= 2 (count (second %))) bot)]
    (reduce process {:bot (into {} state) :output output} to-do)))

(def steps (iterate step {:bot init :output {}}))

(defn find-comparing-bot
  [steps chips]
  (->> (drop-while (complement (fn [{:keys [bot]}] (seq (filter #(= chips (set (second %))) bot)))) steps)
       first :bot (filter #(= chips (set (second %)))) first first))

(println (find-comparing-bot steps #{61 17}))

(defn find-stable-state
  [steps]
  (first (first (drop-while #(not= (first %) (second %)) (partition 2 1 steps)))))

(let [outputs (:output (find-stable-state steps))]
  (println (reduce * (map (comp first outputs) (range 3)))))
