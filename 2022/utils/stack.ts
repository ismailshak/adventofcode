export default class Stack<T> {
  private storage: Array<T> = [];

  constructor(initial = []) {
    initial.forEach((item) => this.storage.push(item));
  }

  push(item: T) {
    this.storage.push(item);
  }

  pushMany(items: Array<T>) {
    items.forEach((item) => this.storage.push(item));
  }

  pop() {
    if (this.size() === 0) throw "attempted to pop an empty stack";
    return this.storage.pop() as string;
  }

  peek() {
    return this.storage[this.size() - 1];
  }

  size() {
    return this.storage.length;
  }

  toString() {
    return this.storage.toString();
  }
}
