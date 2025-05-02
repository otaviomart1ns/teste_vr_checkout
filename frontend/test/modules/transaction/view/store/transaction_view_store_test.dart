import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/modules/transaction/view/store/transaction_view_store.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:mocktail/mocktail.dart';

class MockTransactionService extends Mock implements TransactionService {}

void main() {
  late TransactionViewStore store;
  late MockTransactionService mockService;

  setUp(() {
    mockService = MockTransactionService();
    store = TransactionViewStore(mockService);
  });

  test('fetchCurrencies deve popular currencies corretamente', () async {
    when(() => mockService.fetchCurrencies())
        .thenAnswer((_) async => ['Brazil-Real', 'Canada-Dollar']);

    await store.fetchCurrencies();

    expect(store.currencies, contains('Brazil-Real'));
    expect(store.currencies, contains('Canada-Dollar'));
    expect(store.currencies.length, 2);
  });

  test('fetchTransaction deve popular transaction corretamente', () async {
    final expected = {
      "id": "a664d78d-cce6-4770-b287-b176a9e6e62a",
      "description": "Moto",
      "date": "2020-01-02",
      "amount_usd": 542.96,
      "exchange_rate": 4.475,
      "amount_converted": 2429.75,
      "to_currency": "Brazil-Real",
      "rate_date": "2019-12-31"
    };

    when(() => mockService.fetchTransaction(
            'a664d78d-cce6-4770-b287-b176a9e6e62a', 'Brazil-Real'))
        .thenAnswer((_) async => expected);

    await store.fetchTransaction(
        'a664d78d-cce6-4770-b287-b176a9e6e62a', 'Brazil-Real');

    expect(store.transaction, isNotNull);
    expect(store.transaction!['description'], 'Moto');
    expect(store.transaction!['amount_converted'], 2429.75);
  });

  test('fetchLatestTransactions deve popular lista corretamente', () async {
    final fakeList = [
      {
        "id": "a664d78d-cce6-4770-b287-b176a9e6e62a",
        "description": "Moto",
        "date": "2020-01-02",
        "amount_usd": 542.96
      }
    ];

    when(() => mockService.fetchLatestTransactions(limit: 5))
        .thenAnswer((_) async => fakeList);

    await store.fetchLatestTransactions(limit: 5);

    expect(store.latestTransactions, isNotEmpty);
    expect(store.latestTransactions.first['description'], 'Moto');
  });
}
