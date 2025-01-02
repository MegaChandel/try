const newTaskInput = document.getElementById('new-task');
const addTaskButton = document.getElementById('add-btn');
const taskList = document.getElementById('task-list');
const pendingCount = document.getElementById('pending-count');

document.addEventListener('DOMContentLoaded', loadTasks);

addTaskButton.addEventListener('click', () => {
  const taskName = newTaskInput.value.trim();
  if (taskName) {
    addTask(taskName);
    saveTasks();
    newTaskInput.value = '';
  }
});

function addTask(name, completed = false) {
  const li = document.createElement('li');
  li.className = `task-item ${completed ? 'completed' : ''}`;
  li.innerHTML = `
    <span>${name}</span>
    <div class="task-actions">
      <button onclick="toggleComplete(this)">✔</button>
      <button onclick="editTask(this)">✎</button>
      <button onclick="deleteTask(this)">✖</button>
    </div>
  `;
  taskList.appendChild(li);
  updatePendingCount();
}

function toggleComplete(button) {
  const task = button.parentElement.parentElement;
  task.classList.toggle('completed');
  saveTasks();
  updatePendingCount();
}

function editTask(button) {
  const task = button.parentElement.parentElement;
  const taskName = prompt('Edit Task', task.children[0].textContent);
  if (taskName !== null) {
    task.children[0].textContent = taskName.trim();
    saveTasks();
  }
}

function deleteTask(button) {
  const task = button.parentElement.parentElement;
  task.remove();
  saveTasks();
  updatePendingCount();
}

function updatePendingCount() {
  const pendingTasks = [...taskList.children].filter(
    (task) => !task.classList.contains('completed')
  ).length;
  pendingCount.textContent = `Pending Tasks: ${pendingTasks}`;
}

function saveTasks() {
  const tasks = [...taskList.children].map((task) => ({
    name: task.children[0].textContent,
    completed: task.classList.contains('completed'),
  }));
  localStorage.setItem('tasks', JSON.stringify(tasks));
}

function loadTasks() {
  const tasks = JSON.parse(localStorage.getItem('tasks')) || [];
  tasks.forEach((task) => addTask(task.name, task.completed));
}
