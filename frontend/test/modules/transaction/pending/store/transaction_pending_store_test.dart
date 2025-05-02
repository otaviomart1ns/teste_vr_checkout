import 'package:flutter_test/flutter_test.dart';
import 'package:frontend/modules/transaction/pending/store/transaction_pending_store.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:mocktail/mocktail.dart';

class MockLocalTransactionService extends Mock
    implements LocalTransactionService {}

class MockTransactionService extends Mock implements TransactionService {}

void main() {
  late TransactionPendingStore store;
  late MockLocalTransactionService mockLocal;
  late MockTransactionService mockRemote;

  final exampleTx = LocalTransaction(
    id: 'a664d78d-cce6-4770-b287-b176a9e6e62a',
    description: 'Moto',
    date: DateTime.parse('2020-01-02'),
    amountUsd: 542.96,
  );

  setUpAll(() {
    registerFallbackValue(exampleTx);
  });

  setUp(() {
    mockLocal = MockLocalTransactionService();
    mockRemote = MockTransactionService();
    store = TransactionPendingStore(mockLocal, mockRemote);
  });

  test('loadPendingTransactions deve popular lista com dados locais', () async {
    when(() => mockLocal.getAllLocalTransactions())
        .thenAnswer((_) async => [exampleTx]);

    await store.loadPendingTransactions();

    expect(store.pendingTransactions.length, 1);
    expect(store.pendingTransactions.first.description, 'Moto');
  });

  test('deletePendingTransaction deve remover item da lista', () async {
    store.pendingTransactions.add(exampleTx);

    when(() => mockLocal.deleteLocalTransaction(any()))
        .thenAnswer((_) async => {});

    await store.deletePendingTransaction(exampleTx.id);

    expect(store.pendingTransactions.isEmpty, isTrue);
  });

  test('editPendingTransaction deve atualizar item', () async {
    store.pendingTransactions.add(exampleTx);

    when(() => mockLocal.saveLocalTransaction(any()))
        .thenAnswer((_) async => {});

    await store.editPendingTransaction(
      id: exampleTx.id,
      newDescription: 'Moto Nova',
      newDate: DateTime(2025),
      newAmountUsd: 999.99,
    );

    final updated = store.pendingTransactions.first;
    expect(updated.description, 'Moto Nova');
    expect(updated.amountUsd, 999.99);
    expect(updated.date.year, 2025);
  });

  test('sendPendingTransaction deve enviar e remover item', () async {
    store.pendingTransactions.add(exampleTx);

    when(() => mockRemote.createTransaction(any())).thenAnswer((_) async => {});
    when(() => mockLocal.deleteLocalTransaction(any()))
        .thenAnswer((_) async => {});

    await store.sendPendingTransaction(exampleTx.id);

    expect(store.pendingTransactions.any((t) => t.id == exampleTx.id), isFalse);
    expect(store.errorMessage, isNull);
  });

  test('sendPendingTransaction com erro deve preencher errorMessage', () async {
    store.pendingTransactions.add(exampleTx);

    when(() => mockRemote.createTransaction(any()))
        .thenThrow(Exception('Falha ao enviar'));

    await store.sendPendingTransaction(exampleTx.id);

    expect(store.errorMessage, contains('Erro ao enviar transação'));
    expect(store.pendingTransactions.length, 1);
  });
}
