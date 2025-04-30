import 'package:mobx/mobx.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';

part 'transaction_view_store.g.dart';

class TransactionViewStore = _TransactionViewStoreBase
    with _$TransactionViewStore;

abstract class _TransactionViewStoreBase with Store {
  final TransactionService _transactionService;

  _TransactionViewStoreBase(this._transactionService);

  @observable
  bool isLoading = false;

  @observable
  Map<String, dynamic>? transaction;

  @observable
  List<String> currencies = [];

  @observable
  List<Map<String, dynamic>> latestTransactions = [];

  @action
  Future<void> fetchCurrencies() async {
    isLoading = true;
    try {
      currencies = await _transactionService.fetchCurrencies();
    } catch (e) {
      currencies = [];
    } finally {
      isLoading = false;
    }
  }

  @action
  Future<void> fetchTransaction(String id, String currency) async {
    isLoading = true;
    transaction = null;
    try {
      transaction = await _transactionService.fetchTransaction(id, currency);
    } catch (e) {
      transaction = null;
    } finally {
      isLoading = false;
    }
  }

  @action
  Future<void> fetchLatestTransactions({int limit = 5}) async {
    isLoading = true;

    try {
      latestTransactions = await _transactionService.fetchLatestTransactions(
        limit: limit,
      );
    } catch (e) {
      latestTransactions = [];
    } finally {
      isLoading = false;
    }
  }
}
