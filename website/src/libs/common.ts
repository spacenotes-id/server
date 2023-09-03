import type { ClassValue } from "clsx";
import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function tw(...classNames: ClassValue[]) {
  return twMerge(clsx(...classNames));
}

type ComposeFunction<T> = (v: T) => T;
export function compose<TData>(fns: ComposeFunction<TData>[]) {
  return (value: TData) => fns.reduce((acc, fun) => fun(acc), value);
}
