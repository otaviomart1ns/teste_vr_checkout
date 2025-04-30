import 'package:mobx/mobx.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';

part 'transaction_pending_store.g.dart';

class TransactionPendingStore = _TransactionPendingStoreBase
    with _$TransactionPendingStore;

abstract class _TransactionPendingStoreBase with Store {
  final LocalTransactionService _localTransactionService;
  final TransactionService _transactionService;

  _TransactionPendingStoreBase(
    this._localTransactionService,
    this._transactionService,
  );

  @observable
  ObservableList<LocalTransaction> pendingTransactions =
      ObservableList<LocalTransaction>();

  @observable
  bool isLoading = false;

  @observable
  String? errorMessage;

  @action
  Future<void> loadPendingTransactions() async {
    isLoading = true;
    errorMessage = null;

    try {
      final transactions = await _localTransactionService
          .getAllLocalTransactions();
      pendingTransactions = ObservableList.of(transactions);
    } catch (e) {
      errorMessage = 'Erro ao carregar transações locais: $e';
    } finally {
      isLoading = false;
    }
  }

  @action
  Future<void> deletePendingTransaction(String id) async {
    try {
      await _localTransactionService.deleteLocalTransaction(id);
      pendingTransactions.removeWhere((t) => t.id == id);
    } catch (e) {
      errorMessage = 'Erro ao deletar transação: $e';
    }
  }

  @action
  Future<void> editPendingTransaction({
    required String id,
    required String newDescription,
    required DateTime newDate,
    required double newAmountUsd,
  }) async {
    try {
      final index = pendingTransactions.indexWhere((t) => t.id == id);
      if (index == -1) {
        throw Exception('Transação não encontrada.');
      }

      final updatedTransaction = pendingTransactions[index].copyWith(
        description: newDescription,
        date: newDate,
        amountUsd: newAmountUsd,
      );
      await _localTransactionService.saveLocalTransaction(updatedTransaction);

      pendingTransactions[index] = updatedTransaction;
    } catch (e) {
      errorMessage = 'Erro ao editar transação: $e';
    }
  }

  @action
  Future<void> sendPendingTransaction(String id) async {
    isLoading = true;
    errorMessage = null;

    try {
      final transaction = pendingTransactions.firstWhere((t) => t.id == id);

      final payload = {
        'description': transaction.description,
        'date': transaction.date.toIso8601String().split('T').first,
        'amount_usd': transaction.amountUsd,
      };

      await _transactionService.createTransaction(payload);

      await _localTransactionService.deleteLocalTransaction(id);
      pendingTransactions.removeWhere((t) => t.id == id);
    } catch (e) {
      errorMessage = 'Erro ao enviar transação: $e';
    } finally {
      isLoading = false;
    }
  }
}
