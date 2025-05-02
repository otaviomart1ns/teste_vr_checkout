import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:dio/dio.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:mocktail/mocktail.dart';

class MockDio extends Mock implements Dio {}

void main() {
  late TransactionService service;
  late MockDio mockDio;

  setUpAll(() async {
    await dotenv.load();
  });

  setUp(() {
    mockDio = MockDio();
    service = TransactionService(dio: mockDio);
  });

  test('fetchCurrencies deve retornar lista de moedas', () async {
    when(() => mockDio.get('/currencies')).thenAnswer(
      (_) async => Response(
        data: ['Canada-Dollar', 'Brazil-Real'],
        statusCode: 200,
        requestOptions: RequestOptions(path: ''),
      ),
    );

    final result = await service.fetchCurrencies();

    expect(result, equals(['Canada-Dollar', 'Brazil-Real']));
  });

  test('fetchTransaction deve retornar dados da transação', () async {
    final expected = {
      'description': 'Compra',
      'amount_usd': 100,
      'amount_converted': 500,
      'exchange_rate': 5,
      'date': '2024-05-01',
    };

    when(() => mockDio.get(
          '/transactions/123',
          queryParameters: {'currency': 'Brazil-Real'},
        )).thenAnswer(
      (_) async => Response(
        data: expected,
        statusCode: 200,
        requestOptions: RequestOptions(path: ''),
      ),
    );

    final result = await service.fetchTransaction('123', 'Brazil-Real');
    expect(result, equals(expected));
  });

  test('fetchLatestTransactions deve retornar lista de transações', () async {
    final data = [
      {'id': '1', 'description': 'A', 'amount_usd': 10, 'date': '2024-01-01'},
    ];

    when(() => mockDio.get(
          '/transactions/latest',
          queryParameters: {'limit': 5},
        )).thenAnswer(
      (_) async => Response(
        data: data,
        statusCode: 200,
        requestOptions: RequestOptions(path: ''),
      ),
    );

    final result = await service.fetchLatestTransactions();
    expect(result, isA<List<Map<String, dynamic>>>());
    expect(result.first['id'], '1');
  });
}
