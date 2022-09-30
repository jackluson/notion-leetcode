/*
 * @lc app=leetcode.cn id=621 lang=javascript
 *
 * [621] 任务调度器
 *
 * https://leetcode.cn/problems/task-scheduler/description/
 *
 * algorithms
 * Medium (58.89%)
 * Likes:    1026
 * Dislikes: 0
 * Total Accepted:    119.1K
 * Total Submissions: 201.6K
 * Testcase Example:  '["A","A","A","B","B","B"]\n2'
 *
 * 给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1
 * 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
 *
 * 然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
 *
 * 你需要计算完成所有任务所需要的 最短时间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tasks = ["A","A","A","B","B","B"], n = 2
 * 输出：8
 * 解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
 * ⁠    在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
 *
 * 示例 2：
 *
 *
 * 输入：tasks = ["A","A","A","B","B","B"], n = 0
 * 输出：6
 * 解释：在这种情况下，任何大小为 6 的排列都可以满足要求，因为 n = 0
 * ["A","A","A","B","B","B"]
 * ["A","B","A","B","A","B"]
 * ["B","B","B","A","A","A"]
 * ...
 * 诸如此类
 *
 *
 * 示例 3：
 *
 *
 * 输入：tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
 * 输出：16
 * 解释：一种可能的解决方案是：
 * ⁠    A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> (待命) -> (待命) -> A ->
 * (待命) -> (待命) -> A
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * tasks[i] 是大写英文字母
 * n 的取值范围为 [0, 100]
 *
 *
 */

// @lc code=start
/**
 * @param {character[]} tasks
 * @param {number} n
 * @return {number}
 */
var leastInterval = function (tasks, n) {
  let total = 0;
  let taskMap = new Map();
  for (const task of tasks) {
    let taskCount = taskMap.get(task) || 0;
    taskMap.set(task, taskCount + 1);
  }
  let preTaskMap = new Map();
  while (taskMap.size) {
    let taskSet = taskMap.keys();
    let selectTask = null;
    for (let [task, preInterval] of preTaskMap) {
      preTaskMap.set(task, preInterval - 1);
    }
    total++;
    for (let task of taskSet) {
      const isMeet = preTaskMap.has(task) ? preTaskMap.get(task) < 0 : true;
      if (
        isMeet &&
        (selectTask === null || taskMap.get(selectTask) < taskMap.get(task))
      ) {
        selectTask = task;
      }
    }
    if (selectTask) {
      preTaskMap.set(selectTask, n);
      let taskCount = taskMap.get(selectTask);
      if (taskCount === 1) {
        taskMap.delete(selectTask);
        preTaskMap.delete(selectTask);
      } else {
        taskMap.set(selectTask, taskCount - 1);
      }
    }
  }
  return total;
};
// @lc code=end
const tasks = [
  'A',
  'A',
  'B',
  'C',
  'D',
  'E',
  'F',
  'G',
  'H',
  'I',
  'J',
  'K',
  'L',
  'M',
  'N',
  'O',
  'P',
  'Q',
  'R',
  'S',
  'T',
  'U',
  'V',
  'W',
  'X',
  'Y',
  'Z',
];
const res = leastInterval(tasks, 29);
console.log('🚀 ~ file: 621.任务调度器.js ~ line 123 ~ res', res);
