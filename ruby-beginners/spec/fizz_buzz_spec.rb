require 'rspec'
require '../lib/fizz_buzz'


describe 'Fizz Buzz' do

  def test_fizz_buzz

    example '3の倍数の場合' do
      it 'Fizzを返す'
        expect(fizz_buzz(3)).to eq  'Fizz'
     end

    example '5の倍数の場合' do
      it 'Buzzを返す'
        expect(fizz_buzz(5)).to eq 'Buzz'
     end

     example '15の倍数の場合' do
       it 'FizzBuzzを返す'
         expect(fizz_buzz(15)).to eq 'FizzBuzz'
     end

  end
end
