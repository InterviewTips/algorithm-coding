# coding=utf-8

from threading import Thread, Lock
class Foo:
    def __init__(self):
        self.mute1 = Lock()
        self.mute2 = Lock()
        self.mute3 = Lock()
        self.mute2.acquire() # 锁住
        self.mute3.acquire() # 锁住


    def first(self, printFirst: 'Callable[[], None]') -> None:

        # printFirst() outputs "first". Do not change or remove this line.
        self.mute1.acquire()
        printFirst()
        self.mute1.release()
        self.mute2.release()


    def second(self, printSecond: 'Callable[[], None]') -> None:

        # printSecond() outputs "second". Do not change or remove this line.
        self.mute2.acquire()
        printSecond()
        self.mute2.release()
        self.mute3.release()


    def third(self, printThird: 'Callable[[], None]') -> None:

        # printThird() outputs "third". Do not change or remove this line.
        self.mute3.acquire()
        printThird()
        self.mute3.release()