# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
use Mix.Config

# This configuration is loaded before any dependency and is restricted
# to this project. If another project depends on this project, this
# file won't be loaded nor affect the parent project. For this reason,
# if you want to provide default values for your application for
# 3rd-party users, it should be done in your "mix.exs" file.

# You can configure your application as:
#
#     config :library, key: :value
#
# and access this configuration in your application as:
#
#     Application.get_env(:library, :key)
#
# You can also configure a 3rd-party app:
#
#     config :logger, level: :info
#

# It is also possible to import configuration files, relative to this
# directory. For example, you can emulate configuration per environment
# by uncommenting the line below and defining dev.exs, test.exs and such.
# Configuration from the imported file will override the ones defined
# here (which is why it is important to import them last).
#
#     import_config "#{Mix.env}.exs"

config :library, mapping: %{
  "a" => 0,
  "b" => 1,
  "c" => 2,
  "d" => 3,
  "e" => 4,
  "f" => 5,
  "g" => 6,
  "h" => 7,
  "i" => 8,
  "j" => 9,
  "k" => 10,
  "l" => 11,
  "m" => 12,
  "n" => 13,
  "o" => 14,
  "p" => 15,
  "q" => 16,
  "r" => 17,
  "s" => 18,
  "t" => 19,
  "u" => 20,
  "v" => 21,
  "w" => 22,
  "x" => 23,
  "y" => 24,
  "z" => 25,
  "," => 26,
  "." => 27,
  " " => 28
}
config :library, inverse_mapping: %{
  0 => "a",
  1 => "b",
  2 => "c",
  3 => "d",
  4 => "e",
  5 => "f",
  6 => "g",
  7 => "h",
  8 => "i",
  9 => "j",
  10 => "k",
  11 => "l",
  12 => "m",
  13 => "n",
  14 => "o",
  15 => "p",
  16 => "q",
  17 => "r",
  18 => "s",
  19 => "t",
  20 => "u",
  21 => "v",
  22 => "w",
  23 => "x",
  24 => "y",
  25 => "z",
  26 => ",",
  27 => ".",
  28 => " "
}

config :library, offset: List.to_tuple(Enum.map(1..10, fn(_) -> Enum.random(1..27) end))

# Library.encode(["a","a","a","a","a","a","a","a","a","a"])
