// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'transaction_create_store.dart';

// **************************************************************************
// StoreGenerator
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, unnecessary_brace_in_string_interps, unnecessary_lambdas, prefer_expression_function_bodies, lines_longer_than_80_chars, avoid_as, avoid_annotating_with_dynamic, no_leading_underscores_for_local_identifiers

mixin _$TransactionCreateStore on _TransactionCreateStoreBase, Store {
  late final _$isLoadingAtom =
      Atom(name: '_TransactionCreateStoreBase.isLoading', context: context);

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
      Atom(name: '_TransactionCreateStoreBase.errorMessage', context: context);

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

  late final _$createTransactionRemoteAsyncAction = AsyncAction(
      '_TransactionCreateStoreBase.createTransactionRemote',
      context: context);

  @override
  Future<void> createTransactionRemote(
      {required String description,
      required DateTime date,
      required double amountUsd}) {
    return _$createTransactionRemoteAsyncAction.run(() => super
        .createTransactionRemote(
            description: description, date: date, amountUsd: amountUsd));
  }

  late final _$createTransactionLocalAsyncAction = AsyncAction(
      '_TransactionCreateStoreBase.createTransactionLocal',
      context: context);

  @override
  Future<void> createTransactionLocal(
      {required String description,
      required DateTime date,
      required double amountUsd}) {
    return _$createTransactionLocalAsyncAction.run(() => super
        .createTransactionLocal(
            description: description, date: date, amountUsd: amountUsd));
  }

  @override
  String toString() {
    return '''
isLoading: ${isLoading},
errorMessage: ${errorMessage}
    ''';
  }
}
