use strict;
use warnings;
use feature qw/say/;

chomp(my $n = <STDIN>);
chomp(my $s = <STDIN>);
my @s_arr = split //, $s;

my $checker = +{
  'A' => 0,
  'B' => 0,
  'C' => 0
};

for (0..$n-1) {
  ++$checker->{$s_arr[$_]};
  if($checker->{'A'} > 0 && $checker->{'B'} > 0 && $checker->{'C'} > 0) {
    say $_ + 1;
    last;
  }
}