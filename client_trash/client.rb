#!/usr/bin/ruby

#require 'grpc'
#require 'calculator_services_pb'

def main
  #create service
  stub = Calculator::Calculator::Stub.new('pantova.digital:50051', :this_channel_is_insecure)
  value = ARGV.size > 0 ?  ARGV[0] : 16
  #trigger message
  number = stub.say_hello(Calculator::Number.new(value: value)).message
  p "#{number}"
end

main()
