defmodule P do
    def fib(0), do: 0
    def fib(1), do: 1
    def fib(n) do
        fib(1, 1, n - 2)
    end

    defp fib(prev, second_prev, 0), do: prev
    defp fib(prev, second_prev, n) do
        new = prev + second_prev
        fib(new, prev, n - 1)
    end

    def fib2(0), do: 0
    def fib2(1), do: 1
    def fib2(2), do: 1
    def fib2(n) do
        fib2(n-1) + fib2(n-2)
    end

    def rev(str) do
        String.split(str, "", trim: true)
        |> rev()
    end

    def rev
end
