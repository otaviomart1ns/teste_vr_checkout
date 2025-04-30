import 'package:dio/dio.dart';
import 'package:frontend/modules/transaction/shared/local_transaction_service.dart';
import 'package:uuid/uuid.dart';
import 'package:flutter/foundation.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:frontend/modules/transaction/shared/transaction_service.dart';
import 'package:mobx/mobx.dart';

part 'transaction_create_store.g.dart';

class TransactionCreateStore = _TransactionCreateStoreBase
    with _$TransactionCreateStore;

abstract class _TransactionCreateStoreBase with Store {
  final TransactionService _transactionService;
  final LocalTransactionService _localTransactionService;

  _TransactionCreateStoreBase(
    this._transactionService,
    this._localTransactionService,
  );

  @observable
  bool isLoading = false;

  @observable
  String? errorMessage;

  @action
  Future<void> createTransactionRemote({
    required String description,
    required DateTime date,
    required double amountUsd,
  }) async {
    isLoading = true;
    errorMessage = null;

    try {
      final payload = {
        'description': description.trim(),
        'date': date.toIso8601String().split('T').first,
        'amount_usd': amountUsd,
      };

      await compute<SendTransactionParams, void>(
        _sendTransactionInBackground,
        SendTransactionParams(
          apiUrl: _transactionService.apiUrl,
          payload: payload,
        ),
      );
    } on DioException catch (e) {
      errorMessage =
          e.response?.data.toString() ??
          e.message ??
          'Erro inesperado ao enviar transação.';
    } catch (e) {
      errorMessage = 'Erro inesperado: $e';
    } finally {
      isLoading = false;
    }
  }

  @action
  Future<void> createTransactionLocal({
    required String description,
    required DateTime date,
    required double amountUsd,
  }) async {
    final transaction = LocalTransaction(
      id: const Uuid().v4(),
      description: description.trim(),
      date: date,
      amountUsd: amountUsd,
    );

    await _localTransactionService.saveLocalTransaction(transaction);
  }
}

class SendTransactionParams {
  final String apiUrl;
  final Map<String, dynamic> payload;

  SendTransactionParams({required this.apiUrl, required this.payload});

  Map<String, dynamic> toMap() {
    return {'apiUrl': apiUrl, 'payload': payload};
  }

  factory SendTransactionParams.fromMap(Map<String, dynamic> map) {
    return SendTransactionParams(
      apiUrl: map['apiUrl'] as String,
      payload: Map<String, dynamic>.from(map['payload'] as Map),
    );
  }
}

Future<void> _sendTransactionInBackground(SendTransactionParams params) async {
  final dio = Dio(BaseOptions(baseUrl: params.apiUrl));
  await dio.post('/transactions', data: params.payload);
}
