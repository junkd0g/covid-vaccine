#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;
use JSON;

my $file = $ARGV[0] or die "Need to get CSV file on the command line\n";
     
open(my $file_data, '<', $file) or die "Could not open '$file' $!\n";

my $data = {};
$data->{data} = [];
my $latest = '';
my $obj = {};
my $counter = 0;
my $total = 0;
while (my $line = <$file_data>) {
    $total++;
    chomp $line;

    my @fields = split "," , $line;

    my $country = lc($fields[2]);

    if (!$latest){
        $latest = $country
    }

    if  ($latest ne $country){
        my $arr = $data->{data};
        push (@$arr,$obj);
        $data->{data} = $arr ;
        $obj = {};
        $counter++;
        $latest = $country
    }
    my $tests = $fields[33];
    my $total_vaccinations = $fields[34];
    my $people_vaccinations = $fields[35];
    my $people_fully_vaccinations = $fields[36];
    my $total_booster = $fields[37];

    my $total_vaccinations_per_hundred = $fields[40];
    my $people_vaccinated_per_hundred = $fields[41];
    my $people_fully_vaccinated_per_hundred = $fields[42];
    my $total_boosters_per_hundred = $fields[43]; 

    #$obj->{$country}->{tests} = add_field($obj->{$country}->{tests},$tests,"test");
    $obj->{name} = $country;

    $obj->{total_vaccinations} = add_field($obj->{total_vaccinations}, $total_vaccinations,"total_vaccinations");
    
    $obj->{people_vaccinations} = add_field($obj->{people_vaccinations}, $people_vaccinations, "people_vaccinations");
    
    $obj->{people_fully_vaccinations} = add_field($obj->{people_fully_vaccinations},$people_fully_vaccinations,"people_fully_vaccinations");

    $obj->{total_booster} = add_field($obj->{total_booster},$total_booster,"total_booster");

    $obj->{total_vaccinations_per_hundred} = $total_vaccinations_per_hundred;
    $obj->{people_vaccinated_per_hundred} = $people_vaccinated_per_hundred;
    $obj->{people_fully_vaccinated_per_hundred} = $people_fully_vaccinated_per_hundred;
    $obj->{total_boosters_per_hundred} = $total_boosters_per_hundred;
}

sub add_field{
    my ($arr,$field,$type) = @_;

    if ($field) {
        push @$arr, int $field;
    }else{
        push @$arr, 0
    } 
    return $arr;
}


my $json = encode_json $data;

my $jdson = decode_json($json);
open my $fh, ">", "data_out.json";
print $fh encode_json($jdson);
close $fh;

print Dumper $total;
print Dumper $counter;