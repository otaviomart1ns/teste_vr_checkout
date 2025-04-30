// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'transaction_view_store.dart';

// **************************************************************************
// StoreGenerator
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, unnecessary_brace_in_string_interps, unnecessary_lambdas, prefer_expression_function_bodies, lines_longer_than_80_chars, avoid_as, avoid_annotating_with_dynamic, no_leading_underscores_for_local_identifiers

mixin _$TransactionViewStore on _TransactionViewStoreBase, Store {
  late final _$isLoadingAtom =
      Atom(name: '_TransactionViewStoreBase.isLoading', context: context);

  @override
  bool get isLoading {
    _$isLoadingAtom.reportRead();
    return super.isLoading;
  }

  @override
  set isLoading(bool value) {
    _$isLoadingAtom.reportWrite(value, super.isLoading, () {
      super.isLoading = value;
    });
  }

  late final _$transactionAtom =
      Atom(name: '_TransactionViewStoreBase.transaction', context: context);

  @override
  Map<String, dynamic>? get transaction {
    _$transactionAtom.reportRead();
    return super.transaction;
  }

  @override
  set transaction(Map<String, dynamic>? value) {
    _$transactionAtom.reportWrite(value, super.transaction, () {
      super.transaction = value;
    });
  }

  late final _$currenciesAtom =
      Atom(name: '_TransactionViewStoreBase.currencies', context: context);

  @override
  List<String> get currencies {
    _$currenciesAtom.reportRead();
    return super.currencies;
  }

  @override
  set currencies(List<String> value) {
    _$currenciesAtom.reportWrite(value, super.currencies, () {
      super.currencies = value;
    });
  }

  late final _$latestTransactionsAtom = Atom(
      name: '_TransactionViewStoreBase.latestTransactions', context: context);

  @override
  List<Map<String, dynamic>> get latestTransactions {
    _$latestTransactionsAtom.reportRead();
    return super.latestTransactions;
  }

  @override
  set latestTransactions(List<Map<String, dynamic>> value) {
    _$latestTransactionsAtom.reportWrite(value, super.latestTransactions, () {
      super.latestTransactions = value;
    });
  }

  late final _$fetchCurrenciesAsyncAction = AsyncAction(
      '_TransactionViewStoreBase.fetchCurrencies',
      context: context);

  @override
  Future<void> fetchCurrencies() {
    return _$fetchCurrenciesAsyncAction.run(() => super.fetchCurrencies());
  }

  late final _$fetchTransactionAsyncAction = AsyncAction(
      '_TransactionViewStoreBase.fetchTransaction',
      context: context);

  @override
  Future<void> fetchTransaction(String id, String currency) {
    return _$fetchTransactionAsyncAction
        .run(() => super.fetchTransaction(id, currency));
  }

  late final _$fetchLatestTransactionsAsyncAction = AsyncAction(
      '_TransactionViewStoreBase.fetchLatestTransactions',
      context: context);

  @override
  Future<void> fetchLatestTransactions({int limit = 5}) {
    return _$fetchLatestTransactionsAsyncAction
        .run(() => super.fetchLatestTransactions(limit: limit));
  }

  @override
  String toString() {
    return '''
isLoading: ${isLoading},
transaction: ${transaction},
currencies: ${currencies},
latestTransactions: ${latestTransactions}
    ''';
  }
}
