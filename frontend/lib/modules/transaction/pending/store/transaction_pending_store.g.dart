// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'transaction_pending_store.dart';

// **************************************************************************
// StoreGenerator
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, unnecessary_brace_in_string_interps, unnecessary_lambdas, prefer_expression_function_bodies, lines_longer_than_80_chars, avoid_as, avoid_annotating_with_dynamic, no_leading_underscores_for_local_identifiers

mixin _$TransactionPendingStore on _TransactionPendingStoreBase, Store {
  late final _$pendingTransactionsAtom = Atom(
      name: '_TransactionPendingStoreBase.pendingTransactions',
      context: context);

  @override
  ObservableList<LocalTransaction> get pendingTransactions {
    _$pendingTransactionsAtom.reportRead();
    return super.pendingTransactions;
  }

  @override
  set pendingTransactions(ObservableList<LocalTransaction> value) {
    _$pendingTransactionsAtom.reportWrite(value, super.pendingTransactions, () {
      super.pendingTransactions = value;
    });
  }

  late final _$isLoadingAtom =
      Atom(name: '_TransactionPendingStoreBase.isLoading', context: context);

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

  late final _$errorMessageAtom =
      Atom(name: '_TransactionPendingStoreBase.errorMessage', context: context);

  @override
  String? get errorMessage {
    _$errorMessageAtom.reportRead();
    return super.errorMessage;
  }

  @override
  set errorMessage(String? value) {
    _$errorMessageAtom.reportWrite(value, super.errorMessage, () {
      super.errorMessage = value;
    });
  }

  late final _$loadPendingTransactionsAsyncAction = AsyncAction(
      '_TransactionPendingStoreBase.loadPendingTransactions',
      context: context);

  @override
  Future<void> loadPendingTransactions() {
    return _$loadPendingTransactionsAsyncAction
        .run(() => super.loadPendingTransactions());
  }

  late final _$deletePendingTransactionAsyncAction = AsyncAction(
      '_TransactionPendingStoreBase.deletePendingTransaction',
      context: context);

  @override
  Future<void> deletePendingTransaction(String id) {
    return _$deletePendingTransactionAsyncAction
        .run(() => super.deletePendingTransaction(id));
  }

  late final _$editPendingTransactionAsyncAction = AsyncAction(
      '_TransactionPendingStoreBase.editPendingTransaction',
      context: context);

  @override
  Future<void> editPendingTransaction(
      {required String id,
      required String newDescription,
      required DateTime newDate,
      required double newAmountUsd}) {
    return _$editPendingTransactionAsyncAction.run(() => super
        .editPendingTransaction(
            id: id,
            newDescription: newDescription,
            newDate: newDate,
            newAmountUsd: newAmountUsd));
  }

  late final _$sendPendingTransactionAsyncAction = AsyncAction(
      '_TransactionPendingStoreBase.sendPendingTransaction',
      context: context);

  @override
  Future<void> sendPendingTransaction(String id) {
    return _$sendPendingTransactionAsyncAction
        .run(() => super.sendPendingTransaction(id));
  }

  @override
  String toString() {
    return '''
pendingTransactions: ${pendingTransactions},
isLoading: ${isLoading},
errorMessage: ${errorMessage}
    ''';
  }
}
